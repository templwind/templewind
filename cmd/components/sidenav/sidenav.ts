class Sidenav extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Sidenav connected");
	}
}

customElements.define("tw-sidenav", Sidenav);