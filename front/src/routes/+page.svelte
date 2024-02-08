
<script lang="ts">
    import { goto } from "$app/navigation";
    import { onDestroy } from "svelte";

    import { userNickname } from '../store'
    let nickname = '' 

    async function getUserData() {
	    userNickname.set(nickname)
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
<input bind:value={nickname} placeholder="Nickname" type='text'/>
<button type="button" on:click={getUserData}>Confirm</button>
