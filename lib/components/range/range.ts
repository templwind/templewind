import './range.scss';

export class TwRange extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Range connected");
	}
}

customElements.define("tw-range", TwRange);