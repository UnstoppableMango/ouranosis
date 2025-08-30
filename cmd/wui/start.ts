import { $ } from 'bun';
import path from 'node:path';

const root = await $`git rev-parse --show-toplevel`.text().then(x => x.trim());
const wuiDir = path.join(root, 'cmd', 'wui');

await Promise.any([
	$`watchexec -w ${root} -w ${wuiDir} -e go -r -- 'go run . --dev'`,
	$`bun run dev`,
]);
