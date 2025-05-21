<script lang="ts">
	import Footer from '../components/Footer.svelte';
	import Header from '../components/Header.svelte';
</script>

<main
	class="bg-background text-foreground font-handwriting mx-auto flex min-h-screen flex-col p-4 lg:w-[60%] lg:p-0"
>
	<Header />

	<section class="flex flex-grow flex-col gap-4">
		<section
			id="root"
			class="mx-auto flex flex-col items-center justify-between lg:w-[90%] lg:flex-row lg:py-12"
		>
			<div class="flex w-full justify-center lg:hidden">
				<img src="/me-sketch.png" alt="me" class="h-64 p-2" />
			</div>
			<div class="flex w-full flex-col text-center text-lg lg:w-fit lg:text-left lg:text-xl">
				<span class="text-2xl font-black lg:text-4xl">Hey! I'm Aditya Jyoti</span>
				<span class="flex items-center justify-center text-sm lg:justify-start lg:text-lg"
					>breaking things since... <img
						src="/shushing-face.svg"
						alt="sush emoji"
						class="aspect-square w-8 lg:w-10"
					/>
				</span>

				<span
					>A diehard engineer using tech to cope with the struggles of life. Susceptible to doing
					things just because they look "cool"</span
				>

				<div class="relative mt-4 flex justify-center gap-4 text-sm font-semibold lg:justify-start">
					<!-- button 1 -->
					<div class="relative">
						<div
							class="handdrawn-1 bg-background border-foreground absolute left-0 top-0 z-10 flex h-10 w-full items-center justify-center border-2 transition-all duration-300 ease-in-out hover:translate-x-0.5 hover:translate-y-0.5"
						>
							<a
								href="https://linkedin.com/in/aditya-jyoti"
								target="_blank"
								class="mt-1 w-full text-center">LINKEDIN</a
							>
						</div>
						<div
							class="handdrawn-1 bg-foreground h-10 w-36 translate-x-1 translate-y-1 lg:w-56"
						></div>
					</div>
					<!-- button 1 -->
					<div class="relative">
						<div
							class="handdrawn-5 text-background bg-foreground absolute left-0 top-0 z-10 flex h-10 w-full items-center justify-center transition-all duration-300 ease-in-out hover:translate-x-0.5 hover:translate-y-0.5"
						>
							<a
								href="https://github.com/aditya-jyoti"
								target="_blank"
								class="handdrawn-3 mt-1 w-full text-center">GITHUB</a
							>
						</div>
						<div
							class="handdrawn-5 bg-background border-foreground h-10 w-36 translate-x-1 translate-y-1 border-2 lg:w-56"
						></div>
					</div>
				</div>
			</div>
			<div class="hidden w-full scale-125 justify-center lg:flex">
				<img src="/me-sketch.png" alt="me" class="h-64 p-2" />
			</div>
		</section>
	</section>

	<Footer />
</main>
