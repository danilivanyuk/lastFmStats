<script lang='ts'>
    import { userNickname } from "../../store";
    import { mockedData as data } from "../../../static/mockData"
    let nickname = $userNickname
    // TODO: if there is no nickname return back to "/"
</script>
<div class="w-[80%] mx-auto">
    <h1>Hello {nickname}</h1>

    <p>Top track: {data.recenttracks.track}</p>

</div>
