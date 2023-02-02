package com.crossasyst.activities.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Column;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ActivityType {
    private String activityTypeCd;
    private String description;
    private Integer activeBit;
}
