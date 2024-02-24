<script lang='ts'>
    import { userNickname } from "../../store";
    import type { TrackType } from "../../core/types/trackTypes";
    import type { IScrobblesDataSet } from "../../core/types/dataSetTypes";
    import { mockedData as globalData } from "../../../static/mockData"
    let nickname = $userNickname
    const data = globalData.recenttracks.track[0] as unknown as TrackType
    let scrobbles: IScrobblesDataSet[] = []
    // TODO: if there is no nickname return back to "/"
    const addArtist = () => {
        for (const [key, value] of Object.entries(globalData.recenttracks.track)) {
            console.log(key, value)
            scrobbles.push({
                track: value.name,
                artist: value.artist["#text"],
                album: value.album["#text"],
                date: parseInt(value.date?.uts || '0') ?? 0
            })
        }
        console.log(scrobbles)
        scrobbles = scrobbles
        // let artistData: IArtistDataSet = Object.entries(data as unknown as TrackType).reduce()
    }
</script>
<div class="w-[80%] mx-auto">
    <h1>Hello {nickname}</h1>
    <button on:click={addArtist}>Load info</button>
    <!-- {uniqueArtist} -->
    {#each (scrobbles) as scrobble}
        <div>
<!-- <img src={data.recenttracks.track[index].image[0]["#text"]} alt=""> -->
            <!-- {data.recenttracks.track[index].image[0]["#text"]} -->
            <p>Artist: {scrobble.artist}</p>
            <p>Track: {scrobble.track}</p>
            <p>Album: {scrobble.album}</p>
            <p>Date: {new Date((scrobble.date)).toISOString()}</p>
            <p>------</p>
        </div>

    {/each}

</div>
