<script>
  import abcjs from 'abcjs';

  let id = 'abc-123';
  const song = 'X:1\nT:Speed the Plough\nM:4/4\nC:Trad.\nK:G\n|:GABc dedB|dedB dedB|c2ec B2dB|c2A2 A2BA|\nGABc dedB|dedB dedB|c2ec B2dB|A2F2 G4:|\n|:g2gf gdBd|g2f2 e2d2|c2ec B2dB|c2A2 A2df|\ng2gf g2Bd|g2f2 e2d2|c2ec B2dB|A2F2 G4:|';
  let data = {
    "Title": { id: null, value: 'title'},
    "Notation": { id: 'notation', value: null},
  };
  setTimeout(function() {abcjs.renderAbc('notation', song, {})}, 300);
</script>

<style>
  .stronger {
    font-weight: 600;
  }

</style>


<div {id} class='stronger'> My ID is {id}</div>

<section>
  <table border="2">
    <tr>
      <th>Field</th>
      <th>Data</th>
    </tr>
    {#each Object.keys(data) as key}
      <tr>
        <td class="field-title">{key}</td>
        <td class="field-value" id={data[key].id}>{@html data[key].value || "[undefined]"}</td>
      </tr>
    {/each}
  </table>
</section>
