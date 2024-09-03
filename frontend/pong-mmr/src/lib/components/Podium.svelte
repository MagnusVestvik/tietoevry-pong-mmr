<script>
// @ts-nocheck

	export let user1;
	export let user2;
	export let user3;

	const placeholderImg = 'https://i.imgur.com/XlLcolo.png';
</script>

<main>
	<div class="bronze">
		<div class="profileimg">
            <img src={user1.imgUrl} alt={user1.name} on:error={(e) => e.target.src = placeholderImg} />
		</div>
		<h1>{user1.name}</h1>
		<h3>{user1.mmr}</h3>
		<div class="pod"></div>
	</div>

	<div class="gold">
		<svg
			class="crown"
			fill="gold"
			viewBox="0 0 640 512"
			height="50px"
			width="50px"
			xmlns="http://www.w3.org/2000/svg"
			><path
				d="M528 448H112c-8.8 0-16 7.2-16 16v32c0 8.8 7.2 16 16 16h416c8.8 0 16-7.2 16-16v-32c0-8.8-7.2-16-16-16zm64-320c-26.5 0-48 21.5-48 48 0 7.1 1.6 13.7 4.4 19.8L476 239.2c-15.4 9.2-35.3 4-44.2-11.6L350.3 85C361 76.2 368 63 368 48c0-26.5-21.5-48-48-48s-48 21.5-48 48c0 15 7 28.2 17.7 37l-81.5 142.6c-8.9 15.6-28.9 20.8-44.2 11.6l-72.3-43.4c2.7-6 4.4-12.7 4.4-19.8 0-26.5-21.5-48-48-48S0 149.5 0 176s21.5 48 48 48c2.6 0 5.2-.4 7.7-.8L128 416h384l72.3-192.8c2.5.4 5.1.8 7.7.8 26.5 0 48-21.5 48-48s-21.5-48-48-48z"
			></path>
		</svg>
		<div class="profileimg">
			<img src={user2.imgUrl} alt={user2.name} on:error={(e) => e.target.src = placeholderImg}    />
		</div>
		<h1>{user2.name}</h1>
		<h3>{user2.mmr}</h3>
		<div class="pod"></div>
	</div>

	<div class="silver">
		<div class="profileimg">
			<img src={user3.imgUrl} alt={user3.name} on:error={(e) => e.target.src = placeholderImg}/>
		</div>
		<h1>{user3.name}</h1>
		<h3>{user3.mmr}</h3>
		<div class="pod"></div>
	</div>
</main>

<style>
	@import url('https://fonts.googleapis.com/css2?family=DM+Sans:ital,opsz,wght@0,9..40,100..1000;1,9..40,100..1000&display=swap');
	main {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		margin: auto;
		width: 60%;
		font-family: 'DM Sans', sans-serif;
	}

	@media (max-width: 1000px) {
		main {
			width: 95%;
		}

		main img {
			width: 5rem;
			height: 5rem;
		}
	}

	@media (min-width: 1001px) {
		main {
			width: 60%;
		}

		main img {
			width: 8rem;
			height: 8rem;
		}
	}

	.profileimg {
		padding: 4px;
		margin: 1rem;
		display: flex;
		align-items: center;
		justify-content: center;
	}
	img {
		object-fit: cover;
		height: 8rem;
		width: 8rem;
		border-radius: 50rem;
	}
	.bronze,
	.silver,
	.gold {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: flex-end;
		text-align: center;
	}
	.bronze .pod {
		width: 100%;
		background: linear-gradient(45deg, #cd853f 0%, #804000 100%);
		height: 2rem;
	}
	.bronze img {
		box-shadow: 0 0 30px rgba(205, 127, 50, 0.5);
	}
	.bronze .profileimg {
		background: radial-gradient(
				ellipse farthest-corner at right bottom,
				#cd853f 0%,
				#b87333 8%,
				#8b5a2b 30%,
				#8a3324 40%,
				transparent 80%
			),
			radial-gradient(
				ellipse farthest-corner at left top,
				#f4a460 0%,
				#cd853f 8%,
				#b87333 25%,
				#804000 62.5%,
				#804000 100%
			);
		border-radius: 50rem;
	}
	.gold .pod {
		width: 100%;
		background: radial-gradient(
				ellipse farthest-corner at right bottom,
				#fedb37 0%,
				#fdb931 8%,
				#9f7928 30%,
				#8a6e2f 40%,
				transparent 80%
			),
			radial-gradient(
				ellipse farthest-corner at left top,
				#ffffff 0%,
				#ffffac 8%,
				#d1b464 25%,
				#5d4a1f 62.5%,
				#5d4a1f 100%
			);
		height: 6rem;
	}
	.gold img {
		box-shadow: 0 0 30px rgba(255, 215, 0, 0.5);
	}
	.gold .profileimg {
		background: linear-gradient(45deg, #ffd700 0%, #b8860a 100%);
		border-radius: 50rem;
	}
	.silver .pod {
		width: 100%;
		background: radial-gradient(
				ellipse farthest-corner at right bottom,
				#c0c0c0 0%,
				#a9a9a9 8%,
				#808080 30%,
				#696969 40%,
				transparent 80%
			),
			radial-gradient(
				ellipse farthest-corner at left top,
				#ffffff 0%,
				#d3d3d3 8%,
				#a9a9a9 25%,
				#808080 62.5%,
				#696969 100%
			);
		height: 4rem;
	}
	.silver img {
		box-shadow: 0 0 30px rgba(212, 212, 212, 0.5);
	}
	.silver .profileimg {
		background: linear-gradient(45deg, white 0%, #333 100%);
		border-radius: 50%;
	}
	.crown {
		transform: translate(-2rem, 1.3rem) rotate(-20deg);
	}
	h1 {
		font-size: 1.7rem;
	}
	h3 {
		margin: 0.5rem;
		color: gray;
	}
</style>
