class AppBar extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("AppBar connected");
	}
}

customElements.define("tw-app-bar", AppBar);