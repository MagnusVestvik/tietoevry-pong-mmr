<script>
	import { onMount } from 'svelte';

	/** @type {boolean} */
	let isHovered = false;

	/** @type {HTMLCanvasElement} */
	let canvas;

	/** @type {CanvasRenderingContext2D} */
	let ctx;

	/**
	 * @typedef {Object} Ball
	 * @property {number} x
	 * @property {number} y
	 * @property {number} radius
	 * @property {number} dx
	 * @property {number} dy
	 */

	/** @type {Ball} */
	let ball;

	/**
	 * @typedef {Object} Paddle
	 * @property {number} x
	 * @property {number} y
	 * @property {number} width
	 * @property {number} height
	 * @property {number} dy
	 */

	/** @type {Paddle} */
	let leftPaddle;

	/** @type {number} */
	let animationId;

	/**
	 * @typedef {Object} Wall
	 * @property {number} x
	 * @property {number} y
	 * @property {number} width
	 * @property {number} height
	 */

	/** @type {Wall} */
	let wall = { x: 290, y: 0, width: 10, height: 150 }; // Static wall

	onMount(() => {
		ctx = /** @type {CanvasRenderingContext2D} */ (canvas.getContext('2d'));
		init();
		animate();
	});

	function init() {
		ball = { x: 150, y: 75, radius: 5, dx: 2, dy: 2 };
		leftPaddle = { x: 10, y: 50, width: 10, height: 50, dy: 2 };
	}

	function animate() {
		animationId = requestAnimationFrame(animate);
		if (isHovered) {
			updateGameState();
		}
		draw();
	}

	function updateGameState() {
		ball.x += ball.dx;
		ball.y += ball.dy;

		// Wall interaction
		if (ball.y - ball.radius < 0 || ball.y + ball.radius > canvas.height) ball.dy *= -1;

		// Bounce of left paddle
		if (
			(ball.x - ball.radius < leftPaddle.x + leftPaddle.width &&
				ball.y > leftPaddle.y &&
				ball.y < leftPaddle.y + leftPaddle.height) ||
			(ball.x + ball.radius > wall.x && ball.y > wall.y && ball.y < wall.y + wall.height)
		)
			ball.dx *= -1;

		// Respawn ball when of bounds
		if (ball.x + ball.radius < 0 || ball.x - ball.radius > canvas.width) init();
	}

	function drawBall() {
		ctx.beginPath();
		ctx.arc(ball.x, ball.y, ball.radius, 0, Math.PI * 2);
		ctx.fillStyle = 'white';
		ctx.fill();
		ctx.closePath();
	}

	function drawPlayerPaddle() {
		ctx.fillRect(leftPaddle.x, leftPaddle.y, leftPaddle.width, leftPaddle.height);
	}

	function drawWall() {
		ctx.fillRect(wall.x, wall.y, wall.width, wall.height);
	}

	function draw() {
		ctx.clearRect(0, 0, canvas.width, canvas.height);
		drawBall();
		drawPlayerPaddle();
		drawWall();
	}

	/**
	 * Handle mouse move event to update left paddle's y position
	 * @param {MouseEvent} event
	 */
	function handleMouseMove(event) {
		let mouseY = event.clientY - canvas.getBoundingClientRect().top;
		if (mouseY > 25 && mouseY < 125) leftPaddle.y = mouseY - 25;
	}
</script>

<canvas
	bind:this={canvas}
	width="300"
	height="150"
	on:mousemove={handleMouseMove}
	on:mouseenter={() => (isHovered = true)}
	on:mouseleave={() => (isHovered = false)}
	style="background: black; cursor: none;"
/>
