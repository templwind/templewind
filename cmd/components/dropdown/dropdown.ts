class Dropdown extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Dropdown connected");
	}
}

customElements.define("tw-dropdown", Dropdown);