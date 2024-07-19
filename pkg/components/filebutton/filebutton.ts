import './filebutton.scss';

export class TwFilebutton extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Filebutton connected");
	}
}

customElements.define("tw-filebutton", TwFilebutton);