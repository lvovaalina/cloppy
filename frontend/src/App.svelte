<script lang="ts">
  import logo from './assets/images/logo-universal.png';
  import {GetClipboardHistory, SetValueToClipboard} from '../wailsjs/go/main/App.js';
	import { onMount } from "svelte";

  let clipboardHistory: string[] = [];

  function greet(): void {
    Greet(name).then(result => resultText = result)
  }

  const updateList:void = () => {
	  GetClipboardHistory().then((result) => {
			clipboardHistory = result;
		});
	};

	const setValueToClipboard: void = (value) => {
		SetValueToClipboard(value).then(() => {
			updateList();
		});
	}

	onMount(() => {
		updateList();
	})
</script>

<main>
	<div class="App">
		<button on:click={updateList}>Refresh</button>

		<ul class="clipboard-history">
			{#each clipboardHistory as text}
				<li on:click={() => (setValueToClipboard(text))} class="clipboard-history-item">{ text }</li>
			{/each}
        </ul>
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