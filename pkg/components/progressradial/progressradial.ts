import './progressradial.scss';

export class TwProgressradial extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Progressradial connected");
	}
}

customElements.define("tw-progressradial", TwProgressradial);