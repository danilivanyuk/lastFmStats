<script lang='ts'>
    import { userNickname } from "../../store";
    import type { TrackType } from "../../core/types/trackTypes";
    import type { IScrobblesDataSet } from "../../core/types/dataSetTypes";
    import { mockedData as globalData } from "../../../static/mockData"
    import * as Table from "$lib/components/ui/table";
    let nickname = $userNickname
    const data = globalData.recenttracks.track[0] as unknown as TrackType
    let scrobbles: IScrobblesDataSet[] = []
    let uniqueArtist: string[] = []
    // TODO: if there is no nickname return back to "/"
    const addArtist = () => {
        for (const [key, value] of Object.entries(globalData.recenttracks.track)) {
            scrobbles.push({
                track: value.name,
                artist: value.artist["#text"],
                album: value.album["#text"],
                date: parseInt(value.date?.uts || '0') ?? 0
            })
        }
        uniqueArtist = Array.from(new Set(scrobbles.map((item: IScrobblesDataSet) => item.artist)))
        scrobbles = scrobbles
    }
    const getArtistScrobbles = (artist: string) => {
        const test = scrobbles.filter(function(scrobble: IScrobblesDataSet) {
        return (scrobble.artist === artist)
        }).length
        return test
    }
    const getUniqueTracks = (artist: string) => {

    }
</script>
<div class="w-[80%] mx-auto">
    <h1>Hello {nickname}</h1>
    <button on:click={addArtist}>Load info</button>
    <Table.Root>
      <Table.Caption>Dataset</Table.Caption>
      <Table.Header>
        <Table.Row>
          <Table.Head class="w-[100px]">Artist</Table.Head>
          <Table.Head>Tracks</Table.Head>
          <Table.Head>Scrobbles</Table.Head>
          <Table.Head class="text-right">Rank</Table.Head>
        </Table.Row>
      </Table.Header>
      <Table.Body>
                {#each (uniqueArtist) as artist}

        <Table.Row>
          <Table.Cell class="font-medium">{artist}</Table.Cell>
          <Table.Cell>{getUniqueTracks(artist)}</Table.Cell>
          <Table.Cell>{getArtistScrobbles(artist)}</Table.Cell>
          <Table.Cell class="text-right">Rank</Table.Cell>
        </Table.Row>
                    {/each}

      </Table.Body>
    </Table.Root>



</div>
