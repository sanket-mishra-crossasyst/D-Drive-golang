package com.crossasyst.activities.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Column;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ProcessingStatusType {
    private String processingStatusTypeCd;
    private String description;
    private Integer activeBit;
}
