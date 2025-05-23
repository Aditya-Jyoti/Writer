<script lang="ts">
	import { goto } from '$app/navigation';

	let username = '';
	let password = '';

	async function login() {
		const res = await fetch('/api/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ username, password })
		});

		const data = await res.json();

		if (!res.ok) {
			alert(data.error || 'Login failed');
			return;
		}

		localStorage.setItem('token', data.token);
		goto('/new');
	}
</script>

<main
	class="bg-background font-handwriting relative flex min-h-screen w-full flex-col items-center justify-center gap-8"
>
	<div class="handdrawn-3 translate-1 absolute z-0 h-[18.5rem] w-[25%] border-2 bg-black"></div>
	<div
		class="handdrawn-1 bg-background hover:translate-0.5 z-10 flex h-[18.5rem] w-[25%] flex-col items-center gap-4 border-2 p-4 transition-transform duration-300 ease-in-out"
	>
		<span class="text-2xl font-semibold underline underline-offset-4"
			>login to continue writing</span
		>

		<div class="flex w-full flex-col items-start justify-start text-lg">
			<label for="username" class="font-semibold">username</label>
			<div class="relative w-full">
				<input disabled class="handdrawn-3 absolute z-0 w-full border-2 bg-black p-1 px-4" />
				<input
					type="text"
					id="username"
					bind:value={username}
					class="handdrawn-4 bg-background hover:translate-0 -translate-0.5 z-10 w-full border-2 p-1 px-4 transition-transform duration-300 ease-in-out"
					placeholder="enter username..."
				/>
			</div>
		</div>
		<div class="flex w-full flex-col items-start justify-start text-lg">
			<label for="password" class="font-semibold">password</label>
			<div class="relative w-full">
				<input disabled class="handdrawn-3 absolute z-0 w-full border-2 bg-black p-1 px-4" />
				<input
					type="password"
					id="password"
					bind:value={password}
					class="handdrawn-4 bg-background hover:translate-0 -translate-0.5 z-10 w-full border-2 p-1 px-4 transition-transform duration-300 ease-in-out"
					placeholder="enter password..."
				/>
			</div>
		</div>

		<div class="relative w-full">
			<button
				disabled
				class="text-background handdrawn-1 bg-background absolute z-0 w-full border-4 border-black p-2 text-lg font-semibold"
				>submit</button
			>
			<button
				class="text-background handdrawn-2 translate-0 hover:translate-0.5 z-10 w-full bg-black p-2 text-lg font-semibold transition-transform duration-300 ease-in-out hover:cursor-pointer"
				on:click={login}>submit</button
			>
		</div>
	</div>
</main>
