import './select.scss';

export class TwSelect extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Select connected");
	}
}

customElements.define("tw-select", TwSelect);