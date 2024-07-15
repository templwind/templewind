class AppRail extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("AppRail connected");
	}
}

customElements.define("tw-app-rail", AppRail);