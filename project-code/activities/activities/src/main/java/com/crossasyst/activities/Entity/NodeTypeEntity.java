package com.crossasyst.activities.Entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Table;

@Entity
@Data
@AllArgsConstructor
@NoArgsConstructor
@Table(name = "node_type")
public class NodeTypeEntity {
    @Column(name = "node_type_cd")

    private String nodeTypeCd;
    @Column(name = "description")

    private String description;
    @Column(name = "active_bit")

    private  Integer activeBit;
}
