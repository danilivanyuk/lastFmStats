
<script lang="ts">
    import { goto } from "$app/navigation";
    import { userNickname } from '../store'
    import Button from "$lib/components/ui/button/button.svelte";
    import * as Card from "$lib/components/ui/card";
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";
    let nickname = '' 

    async function getUserData() {
	    userNickname.set(nickname)
	    goto('/user')
    }
</script>
<div class="grid h-screen place-items-center">
  <Card.Root class="w-[350px] items-center">
    <Card.Header>
      <Card.Title>LastFmStats</Card.Title>
      <Card.Description>Deep dive into yout lastfm stat</Card.Description>
    </Card.Header>
    <Card.Content>
      <form>
        <div class="grid w-full items-center gap-4">
          <div class="flex flex-col space-y-1.5">
            <Label for="name">Enter your lastfm nickname</Label>
            <Input bind:value={nickname} placeholder="Nickname" type='text'/>
          </div>
        </div>
      </form>
    </Card.Content>
    <Card.Footer class="flex justify-center">
      <Button type="button" on:click={getUserData} placeholder={"Confirm"}>Confirm</Button>
    </Card.Footer>
  </Card.Root>
</div>
