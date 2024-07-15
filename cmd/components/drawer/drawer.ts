class Drawer extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Drawer connected");
	}
}

customElements.define("tw-drawer", Drawer);