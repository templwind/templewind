import './breadcrumb.scss';

export class TwBreadcrumb extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Breadcrumb connected");
	}
}

customElements.define("tw-breadcrumb", TwBreadcrumb);