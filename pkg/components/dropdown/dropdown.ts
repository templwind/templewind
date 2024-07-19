import './dropdown.scss';

export class TwDropdown extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Dropdown connected");
	}
}

customElements.define("tw-dropdown", TwDropdown);