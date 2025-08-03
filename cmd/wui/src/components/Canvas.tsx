import { type ReactElement, useEffect, useRef } from 'react';

export interface CanvasProps {
	draw(context: CanvasRenderingContext2D): void;
}

const Canvas = ({ draw }: CanvasProps): ReactElement => {
	const canvasRef = useRef<HTMLCanvasElement>(null);

	useEffect(() => {
		if (canvasRef.current === null) return;

		const canvas = canvasRef.current;
		const context = canvas.getContext('2d');

		if (context === null) return;

		draw(context);
	});

	return <canvas ref={canvasRef} width={200} height={100} />;
};

export default Canvas;
