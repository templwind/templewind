class Navbar extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Navbar connected");
	}
}

customElements.define("tw-navbar", Navbar);