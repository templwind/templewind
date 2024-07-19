import './rail.scss';

export class TwRail extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Rail connected");
	}
}

customElements.define("tw-rail", TwRail);