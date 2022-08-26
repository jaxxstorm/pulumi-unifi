// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.unifi.iam.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetUserPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserPlainArgs Empty = new GetUserPlainArgs();

    /**
     * The MAC address of the user.
     * 
     */
    @Import(name="mac", required=true)
    private String mac;

    /**
     * @return The MAC address of the user.
     * 
     */
    public String mac() {
        return this.mac;
    }

    /**
     * The name of the site the user is associated with.
     * 
     */
    @Import(name="site")
    private @Nullable String site;

    /**
     * @return The name of the site the user is associated with.
     * 
     */
    public Optional<String> site() {
        return Optional.ofNullable(this.site);
    }

    private GetUserPlainArgs() {}

    private GetUserPlainArgs(GetUserPlainArgs $) {
        this.mac = $.mac;
        this.site = $.site;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserPlainArgs $;

        public Builder() {
            $ = new GetUserPlainArgs();
        }

        public Builder(GetUserPlainArgs defaults) {
            $ = new GetUserPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param mac The MAC address of the user.
         * 
         * @return builder
         * 
         */
        public Builder mac(String mac) {
            $.mac = mac;
            return this;
        }

        /**
         * @param site The name of the site the user is associated with.
         * 
         * @return builder
         * 
         */
        public Builder site(@Nullable String site) {
            $.site = site;
            return this;
        }

        public GetUserPlainArgs build() {
            $.mac = Objects.requireNonNull($.mac, "expected parameter 'mac' to be non-null");
            return $;
        }
    }

}
