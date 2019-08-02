<script>
  import abcjs from 'abcjs';
  import { onMount } from 'svelte';

  function Field(id, expander, value) {
    this.id = id;
    this.expander = expander;
    this.value = value;
  }

  function extractField(fieldName) {
    return (s) => {return s[fieldName] || "[not defined]"}
  }

  const DATA_FIELDS = {
    "Title": new Field(null, false, extractField("title")),
    "Notation": new Field("notation", true, (s) => {"[undefined]"}),
    "ABC Notation": new Field(null, true, extractField("abc"))
  };

  const dataObj = {
    title: "the_title",
    abc: 'X:1\nT:Speed the Plough\nM:4/4\nC:Trad.\nK:G\n|:GABc dedB|dedB dedB|c2ec B2dB|c2A2 A2BA|\nGABc dedB|dedB dedB|c2ec B2dB|A2F2 G4:|\n|:g2gf gdBd|g2f2 e2d2|c2ec B2dB|c2A2 A2df|\ng2gf g2Bd|g2f2 e2d2|c2ec B2dB|A2F2 G4:|',
  }

  onMount(() => {
    if(dataObj.abc && dataObj.abc.length > 0) {
      abcjs.renderAbc('notation', dataObj.abc, {});
    }
  })
</script>

<style>
  .stronger {
    font-weight: 600;
  }

</style>

<section>
  <table border="2">
    <tr>
      <th>Field</th>
      <th>Data</th>
    </tr>
    {#each Object.keys(DATA_FIELDS) as key}
      <tr>
        <td class="field-title">{key}</td>
        <td class="field-value" id={DATA_FIELDS[key].id}>{DATA_FIELDS[key].value(dataObj)}</td>
      </tr>
    {/each}
  </table>
</section>
