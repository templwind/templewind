import './heading.scss';

export class TwHeading extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Heading connected");
	}
}

customElements.define("tw-heading", TwHeading);