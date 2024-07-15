class AppShell extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("AppShell connected");
	}
}

customElements.define("tw-app-shell", AppShell);