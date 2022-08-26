// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.unifi.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DevicePortOverride {
    /**
     * @return Human-readable name of the port.
     * 
     */
    private @Nullable String name;
    /**
     * @return Switch port number.
     * 
     */
    private Integer number;
    /**
     * @return ID of the Port Profile used on this port.
     * 
     */
    private @Nullable String portProfileId;

    private DevicePortOverride() {}
    /**
     * @return Human-readable name of the port.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return Switch port number.
     * 
     */
    public Integer number() {
        return this.number;
    }
    /**
     * @return ID of the Port Profile used on this port.
     * 
     */
    public Optional<String> portProfileId() {
        return Optional.ofNullable(this.portProfileId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DevicePortOverride defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private Integer number;
        private @Nullable String portProfileId;
        public Builder() {}
        public Builder(DevicePortOverride defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.number = defaults.number;
    	      this.portProfileId = defaults.portProfileId;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder number(Integer number) {
            this.number = Objects.requireNonNull(number);
            return this;
        }
        @CustomType.Setter
        public Builder portProfileId(@Nullable String portProfileId) {
            this.portProfileId = portProfileId;
            return this;
        }
        public DevicePortOverride build() {
            final var o = new DevicePortOverride();
            o.name = name;
            o.number = number;
            o.portProfileId = portProfileId;
            return o;
        }
    }
}
