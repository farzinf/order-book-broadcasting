import type { WebSocketData } from '$lib/types/WebSocketData';
import { averagePrice } from '$lib/stores/averagePrice';

export class WebSocketService {
	private ws: WebSocket;

	constructor(url: string) {
		this.ws = new WebSocket(url);

		this.ws.onmessage = (event: MessageEvent) => {
			const data: WebSocketData = JSON.parse(event.data);
			averagePrice.set(data.averagePrice);
		};

		this.ws.onopen = () => console.log('Connected to WebSocket');
		this.ws.onclose = () => console.log('Disconnected from WebSocket');
	}

	close() {
		this.ws.close();
	}
}
