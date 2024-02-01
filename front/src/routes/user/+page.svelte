<script lang='ts'>
    import { getContext } from "svelte";

    const user = getContext('user');
    console.log(user)

</script>
<h1>Hello user</h1>
