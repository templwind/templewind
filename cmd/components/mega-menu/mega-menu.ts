class MegaMenu extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("MegaMenu connected");
	}
}

customElements.define("tw-mega-menu", MegaMenu);