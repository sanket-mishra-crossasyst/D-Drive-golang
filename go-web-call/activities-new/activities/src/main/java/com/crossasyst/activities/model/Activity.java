package com.crossasyst.activities.model;

import com.crossasyst.activities.entity.ActivityTypeEntity;
import com.crossasyst.activities.entity.MessageEntity;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Column;
import java.sql.Timestamp;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Activity {
    private Long activityId;
    private String activityName;
    private Timestamp processingStartDate;
    private Timestamp processingEndDate;
    private Long revision;
    private MessageEntity messageEntity;
    private ActivityTypeEntity activityTypeEntity;

}
