import './progress.scss';

export class TwProgress extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Progress connected");
	}
}

customElements.define("tw-progress", TwProgress);