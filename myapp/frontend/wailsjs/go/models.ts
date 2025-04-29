export namespace environment {
	
	export class District {
	    Name: string;
	
	    static createFrom(source: any = {}) {
	        return new District(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	    }
	}
	export class Flight {
	    A: District;
	    B: District;
	    Time_: number;
	
	    static createFrom(source: any = {}) {
	        return new Flight(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.A = this.convertValues(source["A"], District);
	        this.B = this.convertValues(source["B"], District);
	        this.Time_ = source["Time_"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

