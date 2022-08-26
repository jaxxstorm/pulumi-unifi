// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.unifi.setting.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MgmtSshKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final MgmtSshKeyArgs Empty = new MgmtSshKeyArgs();

    /**
     * Comment.
     * 
     */
    @Import(name="comment")
    private @Nullable Output<String> comment;

    /**
     * @return Comment.
     * 
     */
    public Optional<Output<String>> comment() {
        return Optional.ofNullable(this.comment);
    }

    /**
     * Public SSH key.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return Public SSH key.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * Name of SSH key.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of SSH key.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Type of SSH key, e.g. ssh-rsa.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of SSH key, e.g. ssh-rsa.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private MgmtSshKeyArgs() {}

    private MgmtSshKeyArgs(MgmtSshKeyArgs $) {
        this.comment = $.comment;
        this.key = $.key;
        this.name = $.name;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MgmtSshKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MgmtSshKeyArgs $;

        public Builder() {
            $ = new MgmtSshKeyArgs();
        }

        public Builder(MgmtSshKeyArgs defaults) {
            $ = new MgmtSshKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param comment Comment.
         * 
         * @return builder
         * 
         */
        public Builder comment(@Nullable Output<String> comment) {
            $.comment = comment;
            return this;
        }

        /**
         * @param comment Comment.
         * 
         * @return builder
         * 
         */
        public Builder comment(String comment) {
            return comment(Output.of(comment));
        }

        /**
         * @param key Public SSH key.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key Public SSH key.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param name Name of SSH key.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of SSH key.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param type Type of SSH key, e.g. ssh-rsa.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of SSH key, e.g. ssh-rsa.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public MgmtSshKeyArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
