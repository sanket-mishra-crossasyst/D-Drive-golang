package com.crossasyst.activities.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Column;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ObjectRef {
    private Long objectRefId;
    private String objectRef;
    private String msgGuid;
    private Integer revision;

}
