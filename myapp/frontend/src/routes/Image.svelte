<script>
  import Sidebar from "../components/Sidebar.svelte";
  import { onMount } from "svelte";
  import { GetFlights } from "../../wailsjs/go/main/App.js";

  let flights = "";
  let drones = [];
  let nextId = 1;

  function sleep(ms) {
    return new Promise((r) => setTimeout(r, ms));
  }

  function randPct() {
    // 随机 10%~90%
    return (Math.random() * 80 + 0).toFixed(1);
  }

  async function update_flights() {
    while (true) {
      const flight = await GetFlights();
      flights = JSON.stringify(flight, null, 2);

      drones = [
        ...drones,
        {
          id: nextId++,
          sx: -50 + "vw",
          sy: -50 + "vh",
          ex: 50 + "vw",
          ey: 50 + "vh",
        },
      ];

      await sleep(1000);
    }
  }

  onMount(update_flights);

  function removeDrone(id) {
    drones = drones.filter((d) => d.id !== id);
  }
</script>

<div
  class="flex w-full h-full bg-white dark:bg-zinc-800 text-zinc-800 dark:text-white"
>
  <Sidebar />
  <div
    class="relative w-full h-full flex justify-center items-center"
    style="
      background: url('/map2.png') center / 80% auto no-repeat fixed;
    "
  >
    {#each drones as d (d.id)}
      <div
        class="drone"
        on:animationend={() => removeDrone(d.id)}
        style="
          --sx: {d.sx};
          --sy: {d.sy};
          --ex: {d.ex};
          --ey: {d.ey};
        "
      />
    {/each}
  </div>
</div>

<style>
  .drone {
    position: absolute;
    width: 20px;
    height: 20px;
    background: red;
    border-radius: 100%;
    animation: fly 30s linear forwards;
    /* 初始放在 var(--sx),var(--sy) ： */
    /* left: var(--sx);
    top: var(--sy); */
    transform: translate(var(--sx), var(--sy));
    opacity: 100;
  }

  @keyframes fly {
    99% {
      /* left: var(--ex);
      top: var(--ey); */
      transform: translate(var(--ex), var(--ey));
      opacity: 100;
    }

    100% {
      /* left: var(--ex);
      top: var(--ey); */
      transform: translate(var(--ex), var(--ey));
      opacity: 0;
    }
  }
</style>
