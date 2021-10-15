<script lang="ts">
	import List, { Item } from "@smui/list";
	import Button, { Label } from "@smui/button";

	export let clipboardHistory = [];

	const handleRefresh = () => {
		window.backend.getClipboardHistory().then((result) => {
			console.log(result);
			clipboardHistory = result;
		});
	};
</script>

<main>
	<div class="App">
		<Button on:click="{handleRefresh}"><Label>Refresh</Label></Button>

		<List class="clipboard-history">
			{#each clipboardHistory as text}
				<Item class="clipboard-history-item">{ text }</Item>
			{/each}
		</List>
	</div>
</main>

<style>
	* :global(.clipboard-history) {
		color: white;
		list-style-type: none;
		padding: 0;
		border: 1px solid #ddd;
	}

	* :global(.clipboard-history-item) {
		padding: 8px 16px;
  		border-bottom: 1px solid #ddd;
	}

	* :global(.clipboard-history-item:last-child) {
		border-bottom: none;
	}
</style>