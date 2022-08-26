// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.unifi.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class WlanScheduleArgs extends com.pulumi.resources.ResourceArgs {

    public static final WlanScheduleArgs Empty = new WlanScheduleArgs();

    /**
     * Time of day to end the block.
     * 
     */
    @Import(name="blockEnd", required=true)
    private Output<String> blockEnd;

    /**
     * @return Time of day to end the block.
     * 
     */
    public Output<String> blockEnd() {
        return this.blockEnd;
    }

    /**
     * Time of day to start the block.
     * 
     */
    @Import(name="blockStart", required=true)
    private Output<String> blockStart;

    /**
     * @return Time of day to start the block.
     * 
     */
    public Output<String> blockStart() {
        return this.blockStart;
    }

    /**
     * Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
     * 
     */
    @Import(name="dayOfWeek", required=true)
    private Output<String> dayOfWeek;

    /**
     * @return Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
     * 
     */
    public Output<String> dayOfWeek() {
        return this.dayOfWeek;
    }

    private WlanScheduleArgs() {}

    private WlanScheduleArgs(WlanScheduleArgs $) {
        this.blockEnd = $.blockEnd;
        this.blockStart = $.blockStart;
        this.dayOfWeek = $.dayOfWeek;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WlanScheduleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WlanScheduleArgs $;

        public Builder() {
            $ = new WlanScheduleArgs();
        }

        public Builder(WlanScheduleArgs defaults) {
            $ = new WlanScheduleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param blockEnd Time of day to end the block.
         * 
         * @return builder
         * 
         */
        public Builder blockEnd(Output<String> blockEnd) {
            $.blockEnd = blockEnd;
            return this;
        }

        /**
         * @param blockEnd Time of day to end the block.
         * 
         * @return builder
         * 
         */
        public Builder blockEnd(String blockEnd) {
            return blockEnd(Output.of(blockEnd));
        }

        /**
         * @param blockStart Time of day to start the block.
         * 
         * @return builder
         * 
         */
        public Builder blockStart(Output<String> blockStart) {
            $.blockStart = blockStart;
            return this;
        }

        /**
         * @param blockStart Time of day to start the block.
         * 
         * @return builder
         * 
         */
        public Builder blockStart(String blockStart) {
            return blockStart(Output.of(blockStart));
        }

        /**
         * @param dayOfWeek Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
         * 
         * @return builder
         * 
         */
        public Builder dayOfWeek(Output<String> dayOfWeek) {
            $.dayOfWeek = dayOfWeek;
            return this;
        }

        /**
         * @param dayOfWeek Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
         * 
         * @return builder
         * 
         */
        public Builder dayOfWeek(String dayOfWeek) {
            return dayOfWeek(Output.of(dayOfWeek));
        }

        public WlanScheduleArgs build() {
            $.blockEnd = Objects.requireNonNull($.blockEnd, "expected parameter 'blockEnd' to be non-null");
            $.blockStart = Objects.requireNonNull($.blockStart, "expected parameter 'blockStart' to be non-null");
            $.dayOfWeek = Objects.requireNonNull($.dayOfWeek, "expected parameter 'dayOfWeek' to be non-null");
            return $;
        }
    }

}
