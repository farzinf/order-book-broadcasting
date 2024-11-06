<script lang="ts">
    import { onMount } from 'svelte';
    import { WebSocketService } from '$lib/services/WebSocketService';
    import { averagePrice } from '$lib/stores/averagePrice';

    const WEBSOCKET_URL = import.meta.env.VITE_WEBSOCKET_URL || 'ws://localhost:8080/ws';
    let wsService: WebSocketService;

    onMount(() => {
        wsService = new WebSocketService(WEBSOCKET_URL);

        return () => wsService.close();
    });
</script>

<main class="flex flex-col items-center justify-center min-h-screen bg-gray-100">
    <h1 class="text-2xl font-bold text-gray-800 mb-4">Average Order Book Price</h1>
    <div class="p-4 bg-white shadow-md rounded-lg">
        <p class="text-lg text-gray-700">
            {$averagePrice !== null ? `Average Price: ${$averagePrice}` : 'Waiting for data...'}
        </p>
    </div>
</main>