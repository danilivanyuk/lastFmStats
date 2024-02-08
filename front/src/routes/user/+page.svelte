<script lang='ts'>
    import { getContext } from "svelte";
    import { userNickname } from "../../store";
    console.log($userNickname)
    let nickname = $userNickname
    console.log('skdfjlsdf', nickname)
</script>
<h1>Hello {nickname}</h1>
