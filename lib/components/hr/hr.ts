import './hr.scss';

export class TwHr extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Hr connected");
	}
}

customElements.define("tw-hr", TwHr);