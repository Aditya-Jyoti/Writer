<script lang="ts">
	import '../app.css';

	let { children } = $props();
</script>

<main class="bg-background " >
	{@render children()}
</main>
