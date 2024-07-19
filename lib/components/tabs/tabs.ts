import './tabs.scss';

export class TwTabs extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Tabs connected");
	}
}

customElements.define("tw-tabs", TwTabs);