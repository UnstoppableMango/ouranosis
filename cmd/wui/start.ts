import { $ } from 'bun';

await Promise.any([
	$`watchexec -e go -r --wrap-process session -- 'go run . --dev'`,
	$`bun run dev`,
]);
