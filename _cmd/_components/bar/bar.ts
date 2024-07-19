import './bar.scss';

export class TwBar extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Bar connected");
	}
}

customElements.define("tw-bar", TwBar);