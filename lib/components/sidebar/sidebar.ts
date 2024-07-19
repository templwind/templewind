import './sidebar.scss';

export class TwSidebar extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Sidebar connected");
	}
}

customElements.define("tw-sidebar", TwSidebar);