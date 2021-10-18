<script lang="ts">
	import List, { Item } from "@smui/list";
	import Button, { Label } from "@smui/button";
import { onMount } from "svelte";

	export let clipboardHistory = [];

	const updateList = () => {
		window.backend.getClipboardHistory().then((result) => {
			clipboardHistory = result;
		});
	};

	const setValueToClipboard = (value) => {
		window.backend.setValueToClipboard(value).then(() => {
			updateList();
		});
	}

	onMount(() => {
		updateList();
	})
</script>

<main>
	<div class="App">
		<Button on:SMUI:action="{updateList}"><Label>Refresh</Label></Button>

		<List class="clipboard-history">
			{#each clipboardHistory as text}
				<Item on:SMUI:action={() => (setValueToClipboard(text))} class="clipboard-history-item">{ text }</Item>
			{/each}
		</List>
	</div>
</main>

<style>
	* :global(.clipboard-history) {
		list-style-type: none;
		padding: 0;
		border: 1px solid grey;
	}

	* :global(.clipboard-history-item) {
		padding: 8px 16px;
  		border-bottom: 1px solid grey;
	}

	* :global(.clipboard-history-item:last-child) {
		border-bottom: none;
	}
</style>