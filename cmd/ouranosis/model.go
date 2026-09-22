package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/unmango/game/num"

	"github.com/unstoppablemango/ouranosis/pkg/character"
	"github.com/unstoppablemango/ouranosis/pkg/game"
	"github.com/unstoppablemango/ouranosis/pkg/save"
	"github.com/unstoppablemango/ouranosis/pkg/training"
)

type screen int

const (
	screenCreate screen = iota
	screenSheet
	screenTraining
)

var (
	titleStyle = lipgloss.NewStyle().Bold(true)
	dimStyle   = lipgloss.NewStyle().Faint(true)
	errStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	selStyle   = lipgloss.NewStyle().Reverse(true)
)

// tickInterval is how often training credits experience while open.
const tickInterval = time.Second

// awayThreshold is the elapsed time above which a return message is shown.
const awayThreshold = 5 * time.Second

type model struct {
	calc     game.Calculator
	savePath string
	screen   screen
	err      error

	// creation
	name  textinput.Model
	class int

	// sheet
	ch     *character.Character
	cursor int
	levels map[character.Stat]int64

	// training
	session     *training.Session
	stats       trainingStats
	away        num.Number
	pendingAway bool
	ticking     bool
}

type trainingStats struct {
	level      int64
	rate       num.Number
	cost       num.Number
	affordable int64
}

type (
	createdMsg  struct{ ch *character.Character }
	levelsMsg   struct{ levels map[character.Stat]int64 }
	tickMsg     struct{ now time.Time }
	trainingMsg struct {
		gained num.Number
		stats  trainingStats
	}
	errMsg struct{ err error }
)

func newModel(calc game.Calculator, savePath string, ch *character.Character) model {
	ti := textinput.New()
	ti.Placeholder = "name"
	ti.CharLimit = 24
	ti.Focus()
	m := model{calc: calc, savePath: savePath, name: ti, ch: ch, levels: map[character.Stat]int64{}}
	if ch != nil {
		m.screen = screenSheet
	}
	return m
}

func (m model) Init() tea.Cmd {
	if m.screen == screenSheet {
		return m.loadLevels()
	}
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg.err
		return m, nil
	case createdMsg:
		m.ch = msg.ch
		m.screen = screenSheet
		return m, tea.Batch(m.persist(), m.loadLevels())
	case levelsMsg:
		m.levels = msg.levels
		return m, nil
	case tickMsg:
		if m.screen != screenTraining {
			m.ticking = false
			return m, nil
		}
		return m, m.tick(msg.now)
	case trainingMsg:
		m.stats = msg.stats
		if m.pendingAway {
			// The first tick after opening carries the offline progress.
			m.away, m.pendingAway = msg.gained, false
		}
		return m, m.schedule()
	case tea.KeyMsg:
		return m.handleKey(msg)
	}
	if m.screen == screenCreate {
		var cmd tea.Cmd
		m.name, cmd = m.name.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if msg.String() == "ctrl+c" {
		return m, tea.Sequence(m.persist(), tea.Quit)
	}
	m.err = nil
	switch m.screen {
	case screenCreate:
		switch msg.String() {
		case "left", "up":
			m.class = (m.class + len(character.Classes) - 1) % len(character.Classes)
		case "right", "down", "tab":
			m.class = (m.class + 1) % len(character.Classes)
		case "enter":
			if strings.TrimSpace(m.name.Value()) == "" {
				m.err = fmt.Errorf("a name is required")
				return m, nil
			}
			return m, m.create()
		default:
			var cmd tea.Cmd
			m.name, cmd = m.name.Update(msg)
			return m, cmd
		}
	case screenSheet:
		switch msg.String() {
		case "q":
			return m, tea.Sequence(m.persist(), tea.Quit)
		case "up", "k":
			m.cursor = (m.cursor + len(character.Stats) - 1) % len(character.Stats)
		case "down", "j":
			m.cursor = (m.cursor + 1) % len(character.Stats)
		case "enter":
			return m.enterTraining()
		}
	case screenTraining:
		switch msg.String() {
		case "q":
			return m, tea.Sequence(m.persist(), tea.Quit)
		case "esc", "backspace":
			m.screen = screenSheet
			m.session = nil
			return m, tea.Batch(m.persist(), m.loadLevels())
		case "b", "enter":
			return m, m.buy()
		}
	}
	return m, nil
}

func (m model) enterTraining() (tea.Model, tea.Cmd) {
	stat := character.Stats[m.cursor]
	m.session = &training.Session{
		Calc:     m.calc,
		Stat:     stat,
		Base:     m.ch.Base[stat],
		Progress: m.ch.Progress[stat],
	}
	m.screen = screenTraining
	m.stats = trainingStats{}
	m.away = num.Zero()
	m.pendingAway = time.Since(m.session.Progress.LastSeen) > awayThreshold
	if m.ticking {
		return m, nil
	}
	m.ticking = true
	return m, func() tea.Msg { return tickMsg{time.Now()} }
}

func (m model) schedule() tea.Cmd {
	return tea.Tick(tickInterval, func(t time.Time) tea.Msg { return tickMsg{t} })
}

func (m model) create() tea.Cmd {
	name, class := strings.TrimSpace(m.name.Value()), character.Classes[m.class]
	return func() tea.Msg {
		ch, err := character.Create(context.Background(), m.calc, name, class, time.Now())
		if err != nil {
			return errMsg{err}
		}
		return createdMsg{ch}
	}
}

func (m model) loadLevels() tea.Cmd {
	ch := m.ch
	return func() tea.Msg {
		levels := make(map[character.Stat]int64, len(character.Stats))
		for _, s := range character.Stats {
			sess := training.Session{Calc: m.calc, Stat: s, Base: ch.Base[s], Progress: ch.Progress[s]}
			lvl, err := sess.Level(context.Background())
			if err != nil {
				return errMsg{err}
			}
			levels[s] = lvl
		}
		return levelsMsg{levels}
	}
}

func (m model) tick(now time.Time) tea.Cmd {
	sess := m.session
	return func() tea.Msg {
		ctx := context.Background()
		gained, err := sess.Tick(ctx, now)
		if err != nil {
			return errMsg{err}
		}
		stats, err := m.readStats(ctx)
		if err != nil {
			return errMsg{err}
		}
		return trainingMsg{gained, stats}
	}
}

func (m model) buy() tea.Cmd {
	sess := m.session
	return func() tea.Msg {
		ctx := context.Background()
		if err := sess.Buy(ctx); err != nil {
			return errMsg{err}
		}
		stats, err := m.readStats(ctx)
		if err != nil {
			return errMsg{err}
		}
		return trainingMsg{num.Zero(), stats}
	}
}

func (m model) readStats(ctx context.Context) (trainingStats, error) {
	var s trainingStats
	var err error
	if s.level, err = m.session.Level(ctx); err != nil {
		return s, err
	}
	if s.rate, err = m.session.RatePerSecond(ctx); err != nil {
		return s, err
	}
	if s.cost, err = m.session.NextCost(ctx); err != nil {
		return s, err
	}
	if s.affordable, err = m.session.Affordable(ctx); err != nil {
		return s, err
	}
	return s, nil
}

func (m model) persist() tea.Cmd {
	ch := m.ch
	if ch == nil {
		return nil
	}
	return func() tea.Msg {
		if err := save.Save(m.savePath, ch); err != nil {
			return errMsg{err}
		}
		return nil
	}
}

func (m model) View() string {
	var b strings.Builder
	switch m.screen {
	case screenCreate:
		b.WriteString(titleStyle.Render("Who are you?") + "\n\n")
		b.WriteString(m.name.View() + "\n\n")
		for i, c := range character.Classes {
			label := " " + string(c) + " "
			if i == m.class {
				label = selStyle.Render(label)
			}
			b.WriteString(label + " ")
		}
		b.WriteString("\n\n" + dimStyle.Render("type a name, arrows pick a class, enter begins"))
	case screenSheet:
		b.WriteString(titleStyle.Render(fmt.Sprintf("%s the %s", m.ch.Name, m.ch.Class)) + "\n\n")
		for i, s := range character.Stats {
			line := fmt.Sprintf("  %-9s %d", s, m.levels[s])
			if i == m.cursor {
				line = selStyle.Render(line)
			}
			b.WriteString(line + "\n")
		}
		b.WriteString("\n" + dimStyle.Render("enter trains the stat, q saves and quits"))
	case screenTraining:
		p := m.session.Progress
		b.WriteString(titleStyle.Render(fmt.Sprintf("Training %s", m.session.Stat)) + "\n\n")
		if !m.away.IsZero() {
			b.WriteString(fmt.Sprintf("While you were away: +%s experience\n\n", m.away.Short()))
		}
		fmt.Fprintf(&b, "  level      %d\n", m.stats.level)
		fmt.Fprintf(&b, "  experience %s\n", p.Balance().Short())
		fmt.Fprintf(&b, "  rate       %s/s\n", m.stats.rate.Short())
		fmt.Fprintf(&b, "  upgrades   %d\n", p.Upgrades)
		fmt.Fprintf(&b, "  next cost  %s (%d affordable)\n", m.stats.cost.Short(), m.stats.affordable)
		b.WriteString("\n" + dimStyle.Render("b buys an upgrade, esc returns, q saves and quits"))
	}
	if m.err != nil {
		b.WriteString("\n\n" + errStyle.Render(m.err.Error()))
	}
	return b.String() + "\n"
}
