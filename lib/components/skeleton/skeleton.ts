import './skeleton.scss';

export class TwSkeleton extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Skeleton connected");
	}
}

customElements.define("tw-skeleton", TwSkeleton);