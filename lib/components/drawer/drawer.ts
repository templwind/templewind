import './drawer.scss';

export class TwDrawer extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Drawer connected");
	}
}

customElements.define("tw-drawer", TwDrawer);