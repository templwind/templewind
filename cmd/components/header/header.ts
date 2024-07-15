class Header extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Header connected");
	}
}

customElements.define("tw-header", Header);