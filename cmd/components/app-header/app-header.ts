class AppHeader extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("AppHeader connected");
	}
}

customElements.define("tw-app-header", AppHeader);