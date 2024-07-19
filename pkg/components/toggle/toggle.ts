import './toggle.scss';

export class TwToggle extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Toggle connected");
	}
}

customElements.define("tw-toggle", TwToggle);