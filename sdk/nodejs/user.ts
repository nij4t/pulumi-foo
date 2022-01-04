// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * The email address of the Grafana user.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Whether to make user an admin.
     */
    public readonly isAdmin!: pulumi.Output<boolean | undefined>;
    /**
     * The username for the Grafana user.
     */
    public readonly login!: pulumi.Output<string | undefined>;
    /**
     * The display name for the Grafana user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The password for the Grafana user.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The numerical ID of the Grafana user.
     */
    public /*out*/ readonly userId!: pulumi.Output<number>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["isAdmin"] = state ? state.isAdmin : undefined;
            resourceInputs["login"] = state ? state.login : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["isAdmin"] = args ? args.isAdmin : undefined;
            resourceInputs["login"] = args ? args.login : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["userId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * The email address of the Grafana user.
     */
    email?: pulumi.Input<string>;
    /**
     * Whether to make user an admin.
     */
    isAdmin?: pulumi.Input<boolean>;
    /**
     * The username for the Grafana user.
     */
    login?: pulumi.Input<string>;
    /**
     * The display name for the Grafana user.
     */
    name?: pulumi.Input<string>;
    /**
     * The password for the Grafana user.
     */
    password?: pulumi.Input<string>;
    /**
     * The numerical ID of the Grafana user.
     */
    userId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * The email address of the Grafana user.
     */
    email: pulumi.Input<string>;
    /**
     * Whether to make user an admin.
     */
    isAdmin?: pulumi.Input<boolean>;
    /**
     * The username for the Grafana user.
     */
    login?: pulumi.Input<string>;
    /**
     * The display name for the Grafana user.
     */
    name?: pulumi.Input<string>;
    /**
     * The password for the Grafana user.
     */
    password: pulumi.Input<string>;
}