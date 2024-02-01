
<script lang="ts">
    import { goto } from "$app/navigation";
    import { setContext } from "svelte";
    import { writable } from "svelte/store";

	let nickname = ''

	    const user = writable();
	    $: user.set(nickname);

	    // ...and add it to the context for child components to access
	    setContext('user', user);
	async function getUserData() {
		// const res = await fetch('', {
		// 	method: 'POST',
		// 	body: JSON.stringify(nickname)
		// })
		// const response = res.json()
		// if (res.status === 200) {
		// 	goto('/user')
	 //    	}
	 //    	else {
		//     console.log('error', res)
	 //    	}
		goto('/user')
	}
</script>

<h1>LastFmStats</h1>
<p>Please enter your lastFm nickname:</p>
<input bind:value={nickname} placeholder="Nickname" type='text'/>
<button type="button" on:click={getUserData}>Confirm</button>
